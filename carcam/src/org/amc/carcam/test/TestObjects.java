package org.amc.carcam.test;
import static org.amc.carcam.ConfigurationFile.propertyName.COMMAND;
import static org.amc.carcam.ConfigurationFile.propertyName.COMMAND_ARGS;
import static org.amc.carcam.ConfigurationFile.propertyName.FILE_DURATION;
import static org.amc.carcam.ConfigurationFile.propertyName.GPSFILE;
import static org.amc.carcam.ConfigurationFile.propertyName.LOCATION;
import static org.amc.carcam.ConfigurationFile.propertyName.LOGFILE;
import static org.amc.carcam.ConfigurationFile.propertyName.NO_OF_FILES;
import static org.amc.carcam.ConfigurationFile.propertyName.PREFIX;
import static org.amc.carcam.ConfigurationFile.propertyName.SAVERATE;
import static org.amc.carcam.ConfigurationFile.propertyName.SUFFIX;
import static org.mockito.Mockito.*;

import java.io.IOException;
import java.nio.file.DirectoryStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;

import org.amc.carcam.ConfigurationFile;
import org.junit.AfterClass;
import org.junit.BeforeClass;
public class TestObjects
{
	
	/**
	 * 
	 * @return Mock org.amc.carcam.ConfigurationFile 
	 */
	public static ConfigurationFile getMockConfigurationFile()
	{
		ConfigurationFile config=mock(ConfigurationFile.class);
		when(config.getProperty(COMMAND)).thenReturn(getTestDirectory().resolve("raspivid.sh").toString());
		when(config.getProperty(COMMAND_ARGS)).thenReturn("-ex night");
		when(config.getProperty(LOCATION)).thenReturn(getTestDirectory().toString());
		when(config.getProperty(FILE_DURATION)).thenReturn("5");
		when(config.getProperty(LOGFILE)).thenReturn("testlog.log");
		when(config.getProperty(GPSFILE)).thenReturn("test.gps");
		when(config.getProperty(SAVERATE)).thenReturn("2");
		when(config.getProperty(PREFIX)).thenReturn("test_");
		when(config.getProperty(SUFFIX)).thenReturn(".h264");
		when(config.getProperty(NO_OF_FILES)).thenReturn("3");
		return config;
	}

	/**
	 * Create Test Directory
	 * Copy Raspivid.sh to the new directory
	 */
	public static void setUp()
	{
		try 
		{
			if(Files.notExists(getTestDirectory()))
			{
				Files.createDirectory(getTestDirectory());
			}
			Path pseudoExecutable=Paths.get(System.getProperty("user.dir"),"src/org/amc/carcam/test/raspivid.sh");
			Files.copy(pseudoExecutable, Paths.get(getMockConfigurationFile().getProperty(COMMAND)), StandardCopyOption.REPLACE_EXISTING);
			
		} 
		catch (IOException e) 
		{
			throw new RuntimeException(e);
		}
	}

	public static Path getTestDirectory()
	{
		return Paths.get(System.getProperty("user.dir"),"/test");
	}
	
	/**
	 * Remove all files created by the Unit tests
	 */
	public static void tearDown()
	{
		try
		{
			if(Files.exists(getTestDirectory()))
			{
				DirectoryStream<Path> dir=Files.newDirectoryStream(getTestDirectory());
				for(Path p:dir)
				{
					//The Test Log not to be deleted
					if(!p.getFileName().toString().equals(getMockConfigurationFile().getProperty(LOGFILE)))
					{
						Files.deleteIfExists(p);
					}
				}
			}
		}
		catch(IOException ioe)
		{
			throw new RuntimeException(ioe);
		}
	}
}
