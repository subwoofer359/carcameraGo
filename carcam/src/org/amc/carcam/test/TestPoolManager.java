package org.amc.carcam.test;

import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.DirectoryStream;
import java.nio.file.FileSystems;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.amc.carcam.Logger;
import org.amc.carcam.PoolManager;
import org.junit.*;

import static org.junit.Assert.*;
import static org.mockito.Mockito.*;

public class TestPoolManager  
{
	private String prefix="test_prefix";
	private String suffix="test_suffix";
	private int NO_OF_FILES=5;
	Path testDirectory=Paths.get(System.getProperty("user.dir"),"/test");
	
	Logger log;
	PoolManager pool;

	@Before
	public void setUp()
	{
		
		try 
		{
			
			if(Files.exists(testDirectory))
			{
				DirectoryStream<Path> filelist=Files.newDirectoryStream(testDirectory);
				for(Path p:filelist)
				{
					Files.delete(p);
					
				}
				//Files.delete(testDirectory); //TODO Get this working
				
			}
			Thread.sleep(1000);
			if(Files.notExists(testDirectory))
			{
				Files.createDirectory(testDirectory);
			}
			
			log=Logger.getInstance(testDirectory.resolve(Paths.get("../log.log")));
			pool=PoolManager.getInstance(testDirectory, prefix, suffix, NO_OF_FILES);
		} 
		catch (IOException e) 
		{
			throw new RuntimeException(e);
		}
		catch(InterruptedException ie)
		{
			
		}
	}

	@Test
	public void testGetNextFileName()
	{
		for(int i=1;i<1000;i++)
		{
			Path p=pool.getNextFilename();
			String expectedName=String.format("%s_%03d%s", prefix,i,suffix);
			assertEquals(p.getFileName().toString(),expectedName);
			pool.addCompleteFile(p);
		}
	}
	
	@Test
	public void testAddCompletedFile()
	{
		//Create three test files :empty, size < MINIMUM FILE and > MINIMUM File size
		Path testFileEmpty=testDirectory.resolve(pool.getNextFilename());
		Path testFilejustNotEnough=testDirectory.resolve(pool.getNextFilename());
		Path testFile=testDirectory.resolve(pool.getNextFilename());
		
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

}