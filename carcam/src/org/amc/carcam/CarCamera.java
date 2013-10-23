package org.amc.carcam;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Properties;
import static org.amc.carcam.ConfigurationFile.propertyName.*;

public class CarCamera 
{
	private String command;
	//private final String command="touch"; //testing string
	private int duration; //Minutes * (minute) * (1000 milliseconds)
	private Path location; // Directory where to store the video 
	private Path logfile;
	
	//GPS file 
	private Path gpsfile;
	private int saverate; // rate to save GPS information
	
	//Path filename=Paths.get("video.h264"); //replaced by PoolManager.getNextFilename()
	
	private Logger log; // Log file service
	
	private PoolManager poolManager; // PoolManager
	
	private Thread gpslogger; //gpslogger thread @todo a lot of work
	
	private String prefix;
	private String suffix;
	private int number_of_files;
	
	private Path configurationFile=Paths.get("CarCamera.config");
	/**
	 * @param args
	 */
	public CarCamera()
	{	
		// load configuration file
		loadConfigurationFile();
		
		log=Logger.getInstance(logfile);
		
		poolManager=PoolManager.getInstance(location, prefix, suffix, number_of_files);
		
		
		
		try
		{
			gpslogger=new Thread(new GPSLogger(gpsfile, saverate));
			gpslogger.start();
		}
		catch(Exception e)
		{
			log.writeToLog(e.getMessage());
		}
	}
	
	
	//Needs to be tested and check for problems
	private void loadConfigurationFile()
	{
		ConfigurationFile configurefile=new ConfigurationFile(configurationFile);
		command=configurefile.getProperty(COMMAND);
		location=Paths.get(configurefile.getProperty(LOCATION)); // Directory where to store the video 
		logfile=location.resolve(configurefile.getProperty(LOGFILE)); 
		gpsfile=location.resolve(Paths.get(configurefile.getProperty(GPSFILE)));
		prefix=configurefile.getProperty(PREFIX);
		suffix=configurefile.getProperty(SUFFIX);
		
		try
		{
			duration=Integer.parseInt(configurefile.getProperty(FILE_DURATION))*60*1000; //Minutes * (minute) * (1000 milliseconds)
			saverate=Integer.parseInt(configurefile.getProperty(SAVERATE)); // rate to save GPS information
			number_of_files=Integer.parseInt(configurefile.getProperty(NO_OF_FILES));
		}
		catch(NumberFormatException nfe)
		{
			log.writeToLog("Configuration File Parsing Error");
		}
		
	}
	
	public void record()
	{
		Runtime r=Runtime.getRuntime();
		log.writeToLog("Starting Camera command");
		
		Path filename=poolManager.getNextFilename();
		
		String arg=String.format("%s -ew night -t %d -o %s",command,duration,filename);
		
		//String arg=String.format("%s %3$s",command,duration,filename);//testing string
		
		log.writeToLog(arg);
		try
		{
			Process process=r.exec(arg);
			
			BufferedReader reader=new BufferedReader(new InputStreamReader(process.getErrorStream()));
			String error="2";
			while(error!=null)
			{
				error=reader.readLine();
				if(error!=null)
				{
					log.writeToLog(error);
				}
			}
			
			reader.close();
			process.waitFor();
			log.writeToLog("Process ID:"+process.exitValue());
			
			//saved successful recording file in PoolManager
			if(process.exitValue()==0)
			{
				poolManager.addCompleteFile(filename);
			}
		}
		catch(Exception e)
		{
			log.writeToLog(e.getMessage());
		}
		finally
		{
			//closeLog();
		}
	}
	
	public static void main(String[] args) throws InterruptedException 
	{
		CarCamera carCamera=new CarCamera();
		while(true)
		{
			carCamera.record();
			//Thread.sleep(1000);
		}
		
	}

}
