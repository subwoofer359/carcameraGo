package org.amc.carcam;
import java.net.*;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

import java.io.*;



public class GPSLogger implements Runnable{
	
	private Path filename; 			//File to store GPS info
	private final int saverate; 			// save gps data every nth second
	
	private boolean shutdown=false; // stop running the thread
	
	private final int port=2947; 	// port of the GPSD server
	private Socket socket;			// connection to GPSD server
	private InetAddress hostAddress;// Address of GPSD server
	private BufferedReader in;		// reading from the socket
	private PrintWriter out; 		// writing to the socket
	private Pattern pattern; 		// Storing the pattern to retrieve lon,lat and speed values from the info received from the GPSD server
	
	/** GPSD Commands */
	private final String WATCH_STREAM="?WATCH={\"enable\":true,\"json\":true}";
	private final String WATCH_POLL="?WATCH={\"enable\":true};";
	private final String WATCH_CLOSE="?WATCH={\"enable\":false}";
	private final String POLL="?POLL;";
	
	
	public GPSLogger(Path filename,int seconds) throws IOException
	{
		
		this.filename=filename;
		this.saverate=seconds;
		
		this.hostAddress=InetAddress.getByName("127.0.0.1");
		//Create connection to the GPSD Server
		this.socket=new Socket(hostAddress,port);
				//connect input and output streams
		this.out=new PrintWriter(new OutputStreamWriter(socket.getOutputStream()));
		this.in=new BufferedReader(new InputStreamReader(socket.getInputStream()));
			
			//save REGEX pattern to retrieve Longitude, Latitude, Speed and Altitude
		this.pattern =Pattern.compile("\\\"(lon|lat|speed|alt)\\\":-?\\d*.\\d*");
		
		
	}
	
	/**
	 * Send initialisation command to the GPSD server
	 */
	private void gpsinit() 
	{
		
		if(socket!=null )
		{
			try
			{
				
				// Send initialisation command
				//String initCommand=WATCH_STREAM;
				String initCommand=WATCH_POLL;
				out.println(initCommand);
				//out.println(initCommand);
				out.flush();
			}
			catch(Exception e)
			{
				e.printStackTrace();
			}
		}
	}
	
	/**
	 * Set the shutdown boolean
	 */
	public void setShutdown()
	{
		this.shutdown=true;
	}
	
	public void shutdown(boolean wait)
	{
		if(wait)
		{
			setShutdown();
		}
			//Cancel request to monitor GPSD
			out.println(WATCH_CLOSE);
			out.flush();
		
	}
	
	/**
	 * Start to log the GPS data
	 */
	private void saveGPSData()
	{
		
		try(BufferedWriter output=Files.newBufferedWriter(filename, Charset.defaultCharset(), StandardOpenOption.CREATE,StandardOpenOption.WRITE,StandardOpenOption.APPEND))
		{
			Matcher matcher;
			while(!shutdown)
			{
				
				out.println(POLL); // Poll GPSD
				out.flush();
				String input=in.readLine();//Read GPS rate
				//System.out.println(input);
				
				matcher=pattern.matcher(input); // Find required info if possible
				
				boolean found=false; //if no match don't save a new line and flush
				while(matcher.find())
				{
					output.write(matcher.group()+" "); // Save date
					found=true;
				}
				if(found)
				{
					output.newLine();
					output.flush();
				}
				
				//Controls the GPS (Poll and save) to file rate
				try
				{	
					Thread.sleep(saverate*1000);
				}
				catch(InterruptedException e)
				{
					//do nothing
				}
			}
		}
		catch(Exception e)
		{
			e.printStackTrace();
		}
	}
	
	
	@Override
	public void run()
	{
		gpsinit();
		saveGPSData();
	}
	
	
	/**
	 * @param args
	 */
	public static void main(String[] args) throws Exception 
	{
		Path path=Paths.get("/home/adrian/gps.txt");
		GPSLogger logger= new GPSLogger(path,10);
		logger.gpsinit();
		logger.saveGPSData();
	

	}
	
}
