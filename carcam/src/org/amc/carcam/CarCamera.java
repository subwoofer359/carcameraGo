package org.amc.carcam;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.file.Path;
import java.nio.file.Paths;

import static org.amc.carcam.ConfigurationFile.propertyName.*;

/**
 * 
 * @author subwoofer359@gmail.com
 *
 */
public class CarCamera 
{
	
	private String command;
	//private final String command="touch"; //testing string
	private String command_args; //runtime arguments for command
	
	private int duration; //Minutes * (minute) * (1000 milliseconds)
	private Path location; // Directory where to store the video 
	private Path logfile;
	
	//GPS file 
	private Path gpsfile;
	private int saverate; // rate to save GPS information
	
	//Path filename=Paths.get("video.h264"); //replaced by PoolManager.getNextFilename()
	
	private Logger log; // Log file service
	
	private PoolManager poolManager; // PoolManager
	
	private Thread gpslogger; //gpslogger thread 
	
	private String prefix;
	private String suffix;
	private int number_of_files;
	
	//private Path configurationFile;
	private ConfigurationFile configurefile;
	/**
	 * 
	 * @param configFile The file to open for the configuration. Needed for testing
	 */
	public CarCamera(ConfigurationFile configFile)
	{	
		try
		{
			this.configurefile=configFile;
			// load configuration file
			loadConfigurationFile();
		
			log=Logger.getInstance(logfile);
		
			poolManager=PoolManager.getInstance(location, prefix, suffix, number_of_files);
		
		
		
		
			gpslogger=new Thread(new GPSLogger(gpsfile, saverate));
			gpslogger.start();
		}
		catch(NumberFormatException nfe)
		{
			log.writeToLog("Configuration File Parsing Error");
			log.writeToLog("Shutting down");
			log.closeLog();
			System.exit(1);
		}
		catch(IOException ioe)
		{
			//catches IOException thrown by GPSLogger
			log.writeToLog(ioe.getMessage());
		}
		
	}
	
	
	//Needs to be tested and check for problems
	private void loadConfigurationFile()
	{
		
		command=configurefile.getProperty(COMMAND);
		command_args=configurefile.getProperty(COMMAND_ARGS);// Command arguments
		location=Paths.get(configurefile.getProperty(LOCATION)); // Directory where to store the video 
		logfile=location.resolve(configurefile.getProperty(LOGFILE)); 
		gpsfile=location.resolve(Paths.get(configurefile.getProperty(GPSFILE)));
		prefix=configurefile.getProperty(PREFIX);
		suffix=configurefile.getProperty(SUFFIX);
		//Might throw NumberFormatException
		duration=Integer.parseInt(configurefile.getProperty(FILE_DURATION))*60*1000; //Minutes * (minute) * (1000 milliseconds)
		saverate=Integer.parseInt(configurefile.getProperty(SAVERATE)); // rate to save GPS information
		number_of_files=Integer.parseInt(configurefile.getProperty(NO_OF_FILES));
		
		
	}
	
	public void record()
	{
		//Runtime r=Runtime.getRuntime();
		log.writeToLog("Starting Camera command");
		
		Path filename=poolManager.getNextFilename();
		
		String arg=String.format("%s %s -t %d -o %s",command,command_args,duration,filename);
		
		//String arg=String.format("%s %3$s",command,duration,filename);//testing string
		
		log.writeToLog(arg);
		try
		{
			ProcessBuilder pb=new ProcessBuilder(arg);
			//Process process=r.exec(arg);
			pb.redirectErrorStream(true);
			pb.redirectError(logfile.toFile());
			Process process=pb.start();
			//BufferedReader reader=new BufferedReader(new InputStreamReader(process.getErrorStream()));
			
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
		ConfigurationFile config=new ConfigurationFile(Paths.get("CarCamera.config"));// Load Configuration file
		CarCamera carCamera=new CarCamera(config);
		while(true)
		{
			carCamera.record();
			//Thread.sleep(1000);
		}
		
	}

}
