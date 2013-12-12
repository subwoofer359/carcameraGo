package org.amc.carcam;

import java.io.BufferedWriter;
import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;
import java.util.Date;

public class Logger 
{
	//private Path logfile;
	
	private static Logger instance;
	
	private BufferedWriter bufwriter;
	
	private Logger(Path logfile)
	{
		//this.logfile=logfile;
		// do some checking
		
		try
		{
			//log file
			if(!Files.exists(logfile))
			{
				Files.createFile(logfile);
			}
			bufwriter=Files.newBufferedWriter(logfile, Charset.defaultCharset(),StandardOpenOption.APPEND);
		}
		catch(IOException e)
		{
			e.printStackTrace();
		}
	}
		
	public static Logger getInstance(Path logfile)
	{
		if(instance==null)
		{
			instance=new Logger(logfile);
		}
		return instance;
	}

	public void writeToLog(String logEntry)
	{
		try
		{
			if(bufwriter!=null)
			{
				Date d=new Date(System.currentTimeMillis());
				bufwriter.write(String.format("%td/%<tm/%<ty %<tH:%<tM:%<tS :",d));
				bufwriter.write(logEntry);
				bufwriter.newLine();
				bufwriter.flush();
			}
		}
		catch(IOException e)
		{
			e.printStackTrace();
		}
	}
	
	public void closeLog()
	{
		try
		{
			bufwriter.close();
		}
		catch(IOException e)
		{
			//do nothing
		}
	}
}
