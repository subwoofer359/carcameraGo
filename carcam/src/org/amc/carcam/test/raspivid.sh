#!/bin/bash
# A bash script that pretends to be raspivid
# It creates an empty file the size the user specifies with -t (x 2000)
# @author Adrian McLaughlin
# date 10/12/13
args=`getopt -a -l "width:,height:,bitrate:,output:,verbose,timeout:,demo,fps:,framerate:,penc,sh:,sharpness:,co:,contrast:,br:,brightness:,sa:,saturation:,ex:,exposure:" t:o:w:h:b:vt:de $*`

#default value so fallocate doesn't complain when no -t argument is passed
duration=1m;

eval set -- "$args";
while true; do
  case "$1" in
	-t|--timeout)
		shift;
      		if [ -n "$1" ]; then
        		echo "File Duration is $1";
			duration=$1;
        		shift;
      		fi
      		;;
    	-o|--output)
      		shift;
      		if [ -n "$1" ]; then
        		echo "Filename is $1";
			filename=$1;
        		shift;
      		fi
      		;;
	-w|--width)
      		shift;
      		if [ -n "$1" ]; then
        	echo "width is $1";
        	shift;
      		fi
      		;;
	-h|--height)
      		shift;
      		if [ -n "$1" ]; then
        		echo "height is $1";
        		shift;
      		fi
      		;;      
	-b|--bitrate)
      		shift;
      		if [ -n "$1" ]; then
        		echo "bitrate is $1";
        		shift;
      		fi
      		;;
	-v|--verbose)
      		shift;
      		echo "Verbose has been set";
      		;;
	-d|--demo)
      		shift;
        	echo "Demo mode has been set";
      		;;
	--fps|--framerate)
      		shift;
     		if [ -n "$1" ]; then
        		echo "Frame rate is $1";
        		shift;
      		fi
      		;;
	--ex|--exposure)
      		shift;
      		if [ -n "$1" ]; then
        		echo "Exposure is set to $1";
        		shift;
      		fi
      		;;
	--)
      		shift;
      		break;
      		;;
  esac
done
#Multiple by 2000 to approximate the file sizes created by the real raspivid
duration=$[$duration*2000]
result=`fallocate -l $duration $filename`

#‐w,  ‐‐ width     :  Set image  width <size>.  Default  1920  
#‐h,  ‐‐ height    :  Set image  height  <size>.  Default  1080  
#‐b,  ‐‐ bitrate   :  Set bitrate. Use bits per second  (e.g.   10MBits/ s would be ‐ b 10000000)  
#‐o,  ‐‐ output    :  Output  filename  <filename>  (to write to  stdout,  use '‐o  ‐ ')   
#‐v,  ‐‐ verbose   :  Output   verbose information  during  run   
#‐t,  ‐‐ timeout   :  Time (in ms) before  takes picture  and  shuts down.   If  not specified,  set to  5s   
#‐d,  ‐‐ demo      :  Run  a  demo mode (cycle  through range of  camera options, no capture)   
#‐fps,  ‐‐ framerate        :  Specify  the  frames  per second  to  record   
#‐e,  ‐‐ penc      :  Display  preview image  *after*  encoding  (shows  compression artifacts)   
#sh, ‐‐ sharpness :  Set image  sharpness ( ‐100  to  100)   
#‐co, ‐‐ contrast  :  Set image  contrast ( ‐100  to  100)   
#‐br, ‐‐ brightness:  Set image  brightness (0  to   100)   
#‐sa, ‐‐ saturation:  Set image  saturation ( ‐100  to  100)   
#‐ISO,  ‐‐ ISO     :  Set capture ISO   
#‐vs, ‐‐ vstab     :  Turn  on  video stablisation   
#‐ev, ‐‐ ev        :  Set EV compensation (-10 to 10) - range corresponds to  -3 stops to +3 stops  
#‐ex, ‐‐ exposure  :  Set exposure  mode (see  Notes)   
#‐awb,  ‐‐ awb     :  Set AWB mode (see  Notes)   
#‐ifx, ‐‐ imxfx    :  Set image  effect (see  Notes)   
#‐cfx, ‐‐ colfx    :  Set colour  effect (U:V)  
#‐mm,  ‐‐ metering :  Set metering  mode (see  Notes)   
#‐rot,  ‐‐ rotation:  Set image  rotation (0 ‐359)   
#‐hf, ‐‐ hflip     :  Set horizontal flip  
#‐vf, ‐‐ vflip     :  Set vertical  flip 
