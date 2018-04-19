#!/bin/bash
for (( i=1; i<=10; i++ ))
do  
   c=$(($i*10))
   p=$(($i*100))
#   echo "Stations:$c Pizza:$p"
   ./pizzashop -s $c -p $p 
done

