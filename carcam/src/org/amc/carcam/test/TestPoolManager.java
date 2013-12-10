package org.amc.carcam.test;

import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.DirectoryStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.amc.carcam.ConfigurationFile;
import org.amc.carcam.ConfigurationFile.propertyName;
import org.amc.carcam.Logger;
import org.amc.carcam.PoolManager;
import org.junit.*;

import static org.junit.Assert.*;
import static org.mockito.Mockito.*;

public class TestPoolManager  
{
	private String prefix="test_prefix";
	private String suffix="test_suffix";
	private int NO_OF_FILES=50;
	
	Logger log;
	PoolManager pool;


	public TestPoolManager()
	{
	
		log=Logger.getInstance(Paths.get(TestObjects.getMockConfigurationFile().getProperty(propertyName.LOGFILE)));
		pool=PoolManager.getInstance(TestObjects.getTestDirectory(), prefix, suffix, NO_OF_FILES);
		
	}
	
	@BeforeClass
	public static void setUp()
	{
		TestObjects.setUp();
	}
	
	@AfterClass
	public static void tearDown()
	{
		TestObjects.tearDown();
	}
	

	@Test
	public void testGetNextFileName()
	{
		//tos.tearDown();
		//tos.setUp();
		int index=0;//File Index
		try
		{
			DirectoryStream<Path> dir=Files.newDirectoryStream(TestObjects.getTestDirectory(),prefix+"*"+suffix);
			List<Path>files=new ArrayList<>(NO_OF_FILES);
			for(Path p:dir)
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
						index=Integer.parseInt(matcher.group(1));
						index++; //Increment count for next file index
					}
					catch(NumberFormatException nfe)
					{
						//don't stop just log it
						log.writeToLog(nfe.getMessage());
					}
				}
				
			}
			files=null; //Don't need this anymore
		} 
		catch (IOException e)
		{
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		for(int i=1;i<1000;i++)
		{
			Path p=pool.getNextFilename();
			String expectedName=String.format("%s_%03d%s", prefix,(index+i),suffix);
			assertEquals(p.getFileName().toString(),expectedName);
			pool.addCompleteFile(p);
		}
	}
	
	@Test
	public void testAddCompletedFile()
	{
		//Create three test files :empty, size < MINIMUM FILE and > MINIMUM File size
		Path testFileEmpty=TestObjects.getTestDirectory().resolve(pool.getNextFilename());
		Path testFilejustNotEnough=TestObjects.getTestDirectory().resolve(pool.getNextFilename());
		Path testFile=TestObjects.getTestDirectory().resolve(pool.getNextFilename());
		
		//Create content for the files
		List<CharSequence> contents=new ArrayList<>();
		List<CharSequence> contentsNearly=new ArrayList<>();
		List<CharSequence> contentsEnough=new ArrayList<>();
		for(int i=0;i<7000;i++)
		{
			contentsNearly.add("Helofkefkregjrgorojgrgjgjgrrgjorgorgrogrorgoogrjjjjjjjjjjjjjjjjjjjpogrpog" +
									"fepflep[epepepfpepflpefepfleepplelpelppleplplfepleplepl");
		}
		
		contentsEnough.addAll(contentsNearly);
		for(int i=0;i<5000;i++)
		{
			contentsEnough.add("Helofkefkregjrgorojgrgjgjgrrgjorgorgrogrorgoogrjjjjjjjjjjjjjjjjjjjpogrpog" +
									"fepflep[epepepfpepflpefepfleepplelpelppleplplfepleplepl");
		}
		
		
		// Actually create the files on the filesystem
		try {
			Files.write(testFileEmpty, contents, Charset.defaultCharset(),	StandardOpenOption.CREATE_NEW);
			Files.write(testFilejustNotEnough, contentsNearly,Charset.defaultCharset(), StandardOpenOption.CREATE_NEW);
			Files.write(testFile, contentsEnough, Charset.defaultCharset(),	StandardOpenOption.CREATE_NEW);
			
			//Check to see if PoolManager handles them properly
			pool.addCompleteFile(testFile);

			pool.addCompleteFile(testFileEmpty);

			pool.addCompleteFile(testFilejustNotEnough);
			Thread.sleep(2000);// Wait for the PoolManager delete threads to finish
			//Test True
			assertTrue(Files.notExists(testFileEmpty));
			assertTrue(Files.notExists(testFilejustNotEnough));
			assertTrue(Files.exists(testFile));

			//Clean Up
			Files.deleteIfExists(testFile);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		catch(InterruptedException ie)
		{
			
		}

	}
	
	@Test
	public void testNoOfFilesIsControlled()
	{
		ConfigurationFile config=TestObjects.getMockConfigurationFile();
		for(int i=0;i<NO_OF_FILES+3;i++)
		{
			Path file=pool.getNextFilename();
			ProcessBuilder pb=new ProcessBuilder(config.getProperty(propertyName.COMMAND),"-t 300000","-o "+file.toString());
			pb.redirectErrorStream(true);
			pb.redirectOutput(ProcessBuilder.Redirect.appendTo(Paths.get(config.getProperty(propertyName.LOGFILE)).toFile()));
			pb.redirectError(ProcessBuilder.Redirect.appendTo(Paths.get(config.getProperty(propertyName.LOGFILE)).toFile()));
			try
			{
				Process process=pb.start();
				process.waitFor();
				if(process.exitValue()==0)
				{
					pool.addCompleteFile(file);
				}
			} 
			catch (IOException e)
			{
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
			catch(InterruptedException ie)
			{
				//do nothing
			}
		}
		
		//Check how many files there are
		try
		{
			DirectoryStream<Path> dir=Files.newDirectoryStream(TestObjects.getTestDirectory(),prefix+"*"+suffix);
			int actualNoOfFiles=0;
			for(Path p:dir)
			{
				actualNoOfFiles++;
			}
			assertEquals(NO_OF_FILES, actualNoOfFiles);
		} 
		catch (IOException e)
		{
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}

}