package org.amc.carcamera;

import org.amc.carcam.CarCamera;
import org.junit.AfterClass;
import org.junit.BeforeClass;
import org.junit.Test;

public class TestCarCamera 
{
	
	
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
	public void testCarCamera()
	{
		CarCamera camera=new CarCamera(TestObjects.getMockConfigurationFile());		
		for(int i=0;i<100;i++)
		{
			camera.record();
		}
		
	}
}
