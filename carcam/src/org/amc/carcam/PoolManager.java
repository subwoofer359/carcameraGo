package org.amc.carcam;

import java.io.BufferedWriter;
import java.io.IOException;

import java.nio.charset.Charset;
import java.nio.file.*;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.Collections;

import java.util.Date;
import java.util.Iterator;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Controls the files created and stored
 * @author adrian
 *
 */
public class PoolManager 
{
	/**
	 * Number of files to keep
	 */
	private int NUMBER_OF_FILES;
	
	private static final int MINFILESIZE=1024*1000; //THe minimum file size to be accepted
	
	private static PoolManager instance=null;
	
	private String prefix; // File prefix
	private String suffix; // File suffix e.g. ".h264"
	/**
	 * The directory the files stored in
	 */
	private Path directory=null; 
	
	private List<Path> files; // collection of files paths
	
	private Logger log;
	
	private int count=0; // File Index to use
	/**
	 * Constructor
	 * @param directory
	 */
	private PoolManager(Path directory,String prefix,String suffix,int number_of_files)
	{
		this.directory=directory;
		this.NUMBER_OF_FILES=number_of_files;
		this.suffix=suffix;
		this.prefix=prefix;
		//Presumes the Logger object has already been called so no need to pass anything but null
		log=Logger.getInstance(null);
		
		
		//Search for old saved video files and place them in a list
		try(DirectoryStream<Path> ds=Files.newDirectoryStream(directory,"*"+this.suffix);)
		{
			
			files=new ArrayList<>(this.NUMBER_OF_FILES);
			for(Path p:ds)
			{
				files.add(p.toAbsolutePath());
			}
			// Sort alphabetically
			Collections.sort(files);
			
			if(files.size()>0)
			{
				String temp=files.get(files.size()-1).getFileName().toString();
				Pattern pattern=Pattern.compile(this.prefix+"_(\\d+)"+suffix);
				Matcher matcher=pattern.matcher(temp);
				if(matcher.find())
				{
					try
					{
						//System.out.println(matcher.group());
						count=Integer.parseInt(matcher.group(1));
						count++; //Increment count for next file index
					}
					catch(NumberFormatException nfe)
					{
						//don't stop just log it
						log.writeToLog(nfe.getMessage());
					}
				}
			}
			
		}
		catch(IOException e)
		{
			log.writeToLog(e.getMessage());
		}
	}
	
	/**
	 * 
	 * @return PoolManager instance
	 */
	public static PoolManager getInstance(Path directory,String prefix,String suffix,int number_of_files)
	{
		if(instance==null)
		{
			instance=new PoolManager(directory,prefix,suffix,number_of_files); 
		}
		return instance;
	}
	
	/**
	 * 
	 * @return next valid filename for the video
	 */
	public Path getNextFilename()
	{
		String name=String.format("%s_%03d%s",prefix,count,suffix);
		count++; // increment the file index
		Path p=directory.resolve(Paths.get(name)); //new File path for the video
		
		
		
		
		
		//System.out.println(files);
		return p;
	}
	
	/**
	 * When the CarCamera has completed a successful recording it wiil
	 * call this function to store the file;
	 * @param p
	 */
	public void addCompleteFile(Path p)
	{
		//check if List is within NUMBER_OF_FILES constraint
		try
		{
			if(Files.exists(p))
			{
				if(Files.size(p)>MINFILESIZE) // new file has content
				{
					while(files.size()>=NUMBER_OF_FILES)
					{
						Iterator<Path> i=files.iterator();
					
						//remove oldest file from ArrayList
						Path old=i.next();
						//System.out.println("removing "+old);
						i.remove();
					
						//Send file to be deleted
						Delete deleteAction =new Delete(old);
						deleteAction.start();
						
						
					}
					files.add(p);// store absolute path
				}
				else	//File has no content then don't add it to the list and then delete
				{
					Delete deleteAction =new Delete(p);
					deleteAction.start();
					this.count--; //decrement count as file wasn't saved, stop a race on the index 
				}
			}
		}
		catch(IOException ioe)
		{
			log.writeToLog(ioe.getMessage());
		}
	}
	
	/**
	 * Thread to delete files
	 * Hopefully reducing the delay of file Deletion
	 * @author adrian
	 *
	 */
	private static class Delete extends Thread
	{
		private Path p;
		private Logger log;
		/**
		 * 
		 * @param p File to be deleted
		 */
		public Delete(Path p)
		{
			this.p=p;
			log=Logger.getInstance(null);
		}
		
		@Override
		public void run()
		{
			try
			{
				Files.deleteIfExists(p);
				if(Files.notExists(p))
				{
					log.writeToLog(p+" is deleted");
					
				}
				else
				{
					log.writeToLog("Problems deleting "+p);
				}
			}
			catch(IOException ioe)
			{
				log.writeToLog(ioe.getMessage());
				ioe.printStackTrace();
				//Error
			}
		}
	}
	
	//for testing only
	public static void main(String[] args)
	{
		Logger log=Logger.getInstance(Paths.get("/home/adrian/log.log"));
		PoolManager pool=PoolManager.getInstance(Paths.get("/home/adrian"), "video", ".test", 3);	
			try
			{
				int i=0;
				while(i<100)
				{
					Path p=pool.getNextFilename();
					Files.createFile(p);
					
					try(BufferedWriter buf=Files.newBufferedWriter(p, Charset.defaultCharset()))
					{
						for(int j=0;j<1000;j++)
						{
							buf.write("Helofkefkregjrgorojgrgjgjgrrgjorgorgrogrorgoogrjjjjjjjjjjjjjjjjjjjpogrpog" +
									"fepflep[epepepfpepflpefepfleepplelpelppleplplfepleplepl");
						}
					}
					
					
					pool.addCompleteFile(p);
					Thread.sleep(2000);
					i++;
					System.out.println(pool.files);
				}
			}
			catch(IOException ioe)
			{
				ioe.printStackTrace();
			}
			catch(InterruptedException ie)
			{
				ie.printStackTrace();
			}
	
	}
}
	