#set -x

jars=`ls *.jar`
for i in $jars
do
   echo $i
   #jar -tvf $i | grep StaticLoggerBinder 
   jar -tvf $i | grep Driver 
done

