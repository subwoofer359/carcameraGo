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
	private int saverate; 			// save gps data every nth second
	
	private boolean shutdown=false; // stop running the thread
	
	private final int port=2947; 	// port of the GPSD server
	private Socket socket;			// connection to GPSD server
	private InetAddress hostAddress;// Address of GPSD server
	private BufferedReader in;		// reading from the socket
	private PrintWriter out; 		// writing to the socket
	private Pattern pattern; 		// Storing the pattern to retrieve lon,lat and speed values from the info received from the GPSD server
	
	
	
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
				String initCommand="?WATCH={\"enable\":true,\"json\":true}";
				out.println(initCommand);
				out.println(initCommand);
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
	public void shutdown()
	{
		this.shutdown=false;
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
				String input=in.readLine();
				//System.out.println(input);
				
				matcher=pattern.matcher(input);
				
				boolean found=false; //if no match don't save a new line and flush
				while(matcher.find())
				{
					output.write(matcher.group()+" ");
					found=true;
				}
				if(found)
				{
					output.newLine();
					output.flush();
				}
				Thread.sleep(saverate*1000);
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
		GPSLogger logger= new GPSLogger(path,1);
		logger.gpsinit();
		logger.saveGPSData();
	

	}
	
}
