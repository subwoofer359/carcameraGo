package org.amc.carcam.test;
import java.nio.file.Path;
import java.nio.file.Paths;

import org.amc.carcam.ConfigurationFile;
import org.junit.*;

import static org.junit.Assert.*;
import static org.amc.carcam.ConfigurationFile.propertyName;

public class TestConfigurationFile 
{
	private String LOCATION="/location/alpha";
	private int SAVERATE=1000;
	private String LOGFILE="logfile.log";
	private String PREFIX="gnome";
	private String COMMAND_ARGS="-w 22 -e 33 -p 3";
	private int FILE_DURATION=53;
	private int NO_OF_FILES=16;
	private String COMMAND="/usr/bin/touch";
	private String SUFFIX=".mpx";
	private String GPSFILE="gps2.log22";

	
	
	@Test
	public void testConfigurationFile()
	{
		Path configurationFile=Paths.get(System.getProperty("user.dir"),"src/org/amc/carcam/test/TestConfig");
		ConfigurationFile cf =new ConfigurationFile(configurationFile);
		//Location
		assertEquals(cf.getProperty(propertyName.LOCATION),LOCATION);
		assertNotEquals(cf.getProperty(propertyName.LOCATION), "bin");
		
		//SaveRate
		assertEquals(Integer.parseInt(cf.getProperty(propertyName.SAVERATE)),SAVERATE);
		assertNotEquals(Integer.parseInt(cf.getProperty(propertyName.SAVERATE)), 2);
		
		//Prefix
		assertEquals(cf.getProperty(propertyName.PREFIX),PREFIX);
		assertNotEquals(cf.getProperty(propertyName.PREFIX), "bin");
		
		//Command_args
		assertEquals(cf.getProperty(propertyName.COMMAND_ARGS),COMMAND_ARGS);
		assertNotEquals(cf.getProperty(propertyName.COMMAND_ARGS), "bin");
		
		//Logfile
		assertEquals(cf.getProperty(propertyName.LOGFILE),LOGFILE);
		assertNotEquals(cf.getProperty(propertyName.LOGFILE), "bin");
		
		//File Duration
		assertEquals(Integer.parseInt(cf.getProperty(propertyName.FILE_DURATION)),FILE_DURATION);
		assertNotEquals(Integer.parseInt(cf.getProperty(propertyName.FILE_DURATION)), 2);
		
		//No of Files
		assertEquals(Integer.parseInt(cf.getProperty(propertyName.NO_OF_FILES)),NO_OF_FILES);
		assertNotEquals(Integer.parseInt(cf.getProperty(propertyName.NO_OF_FILES)), 2);
		
		//Command
		assertEquals(cf.getProperty(propertyName.COMMAND),COMMAND);
		assertNotEquals(cf.getProperty(propertyName.COMMAND), "bin");
		
		//suffix
		assertEquals(cf.getProperty(propertyName.SUFFIX),SUFFIX);
		assertNotEquals(cf.getProperty(propertyName.SUFFIX), "bin");
		
		//gps
		assertEquals(cf.getProperty(propertyName.GPSFILE),GPSFILE);
		assertNotEquals(cf.getProperty(propertyName.GPSFILE), "bin");
	}
	
	@Test
	public void testSaveConfig()
	{
		Path configurationFile=Paths.get(System.getProperty("user.dir"),"src/org/amc/carcam/test/TestConfig");
		ConfigurationFile cf =new ConfigurationFile(configurationFile);
		
		cf.saveConfigurationInfo();
		testConfigurationFile();
	}
}
