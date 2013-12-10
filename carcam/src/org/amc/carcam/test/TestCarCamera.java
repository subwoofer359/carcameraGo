package org.amc.carcam.test;

import org.amc.carcam.CarCamera;
import org.amc.carcam.ConfigurationFile;

import static org.amc.carcam.ConfigurationFile.propertyName.*;

import org.junit.Test;

import static org.mockito.Mockito.*;

public class TestCarCamera 
{
	ConfigurationFile config;
	
	public TestCarCamera()
	{
		config=mock(ConfigurationFile.class);
		when(config.getProperty(COMMAND)).thenReturn("/bin/touch");
		when(config.getProperty(COMMAND_ARGS)).thenReturn("");
		when(config.getProperty(LOCATION)).thenReturn("/mnt/external");
		when(config.getProperty(FILE_DURATION)).thenReturn("5");
		when(config.getProperty(LOGFILE)).thenReturn("testlog.log");
		when(config.getProperty(GPSFILE)).thenReturn("test.gps");
		when(config.getProperty(SAVERATE)).thenReturn("2");
		when(config.getProperty(PREFIX)).thenReturn("test_");
		when(config.getProperty(SUFFIX)).thenReturn(".h264");
		when(config.getProperty(NO_OF_FILES)).thenReturn("3");
	}
	@Test
	public void testCarCamera()
	{
		CarCamera camera=new CarCamera(this.config);		
	}
}
