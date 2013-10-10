package org.amc.carcam;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.nio.file.Path;
import java.nio.file.Paths;

public class CarCamera 
{
	private final String command="raspivid";
	//private final String command="touch"; //testing string
	private final int duration=5*60*1000; //Minutes * (minute) * (1000 milliseconds)
	private Path location=Paths.get("/mnt/external"); // Directory where to store the video 
	private Path logfile=location.resolve(Paths.get("camera.log"));
	
	//GPS file 
	private Path gpsfile=location.resolve(Paths.get("gps.log"));
	private int saverate=1;
	
	//Path filename=Paths.get("video.h264"); //replaced by PoolManager.getNextFilename()
	
	private Logger log; // Log file service
	
	private PoolManager poolManager; // PoolManager
	
	private Thread gpslogger; //gpslogger thread @todo a lot of work
	
	private String prefix="video";
	private String suffix=".h264";
	private int number_of_files=6;
	
	/**
	 * @param args
	 */
	public CarCamera()
	{
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
	
	
	
	public void record()
	{
		Runtime r=Runtime.getRuntime();
		log.writeToLog("Starting Camera command");
		
		Path filename=poolManager.getNextFilename();
		
		String arg=String.format("%s -t %d -o %s",command,duration,filename);
		
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
