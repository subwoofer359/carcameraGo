package org.amc.carcam;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Properties;

/**
 * Class to handle load and save configuration info to a file
 * Store defaults
 * @author subwoofer359@gmail.com
 *
 */
public class ConfigurationFile 
{

	public enum propertyName
	{
		COMMAND,
		COMMAND_ARGS,
		FILE_DURATION,
		LOCATION,
		LOGFILE,
		GPSFILE,
		SAVERATE,
		PREFIX,
		SUFFIX,
		NO_OF_FILES;
		
	}
	
	private Properties prop;
	
	private Path configurationfile; //File name of the configuration file
	
	private String[][] defaults=
		{
			{propertyName.COMMAND.toString(),"raspivid"},
			{propertyName.COMMAND_ARGS.toString(),""},
			{propertyName.FILE_DURATION.toString(),"5"},
			{propertyName.LOCATION.toString(),"/mnt/external"},
			//{propertyName.LOCATION.toString(),"/home/adrian/external"},
			{propertyName.LOGFILE.toString(),"camera.log"},
			{propertyName.GPSFILE.toString(),"gps.log"},
			{propertyName.SAVERATE.toString(),"30"},
			{propertyName.PREFIX.toString(),"video"},
			{propertyName.SUFFIX.toString(),".h264"},
			{propertyName.NO_OF_FILES.toString(),"6"},
			
		};
	
	
	public ConfigurationFile(Path configurationfile)
	{
		this.configurationfile=configurationfile;
		
		prop=new Properties();
	
		loadConfigurationInfo();
		
	}
	
	private void loadConfigurationInfo()
	{
		if(Files.exists(configurationfile))
		{
			try(BufferedReader reader=Files.newBufferedReader(configurationfile, Charset.defaultCharset()))
			{
				prop.load(reader);
			}
			catch(IOException ioe)
			{
				ioe.printStackTrace();
			}
		}
		else
		{
			loadDefaults();
		}
	}
	
	private void loadDefaults()
	{
		for(String[] a:defaults)
		{
			prop.setProperty(a[0], a[1]);
		}
		saveConfigurationInfo();
	}
	
	private void saveConfigurationInfo()
	{
		try(BufferedWriter writer=Files.newBufferedWriter(configurationfile, Charset.defaultCharset()))
		{
			this.prop.store(writer,"");
		}
		catch(IOException ioe)
		{
			ioe.printStackTrace();
		}
	}
	
	/**
	 * 
	 * @param propertyName
	 * @return if property doesn't exist return empty String not null
	 */
	public String getProperty(propertyName propertyName)
	{
		String result=prop.getProperty(propertyName.toString());
		if(result==null)
		{
			return "";
		}
		else
		{
			return result;
		}
	}
	
	public static void main(String[] args)
	{
		
		ConfigurationFile cfile=new ConfigurationFile(Paths.get("config3.file"));
		System.out.println("Command:"+cfile.getProperty(propertyName.COMMAND));
		
		
	}
	
}
