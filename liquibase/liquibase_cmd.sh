export LIQUIBASE_HOME=/usr/local/Cellar/liquibase/3.10.1/libexec


#Run this command to apply changesets to tinyurlgorm.db.  All changesets are defined in liquibase-changelog.xml file. As an example, two new tables users and accounts are created into the existing db tinyurlgorm.db 

java -Xbootclasspath/a:slf4j-api-1.7.30.jar:logback-classic-1.2.3.jar:logback-core-1.2.3.jar:sqlite-jdbc-3.30.1.jar -classpath sqlite-jdbc-3.30.1.jar:. -jar liquibase.jar --logLevel=debug --driver=org.sqlite.JDBC --url="jdbc:sqlite:../tinyurlgorm.db" --changeLogFile=./liquibase-changelog.xml update

#update history (this requires liquibase.properties file configured)
#liquibase --outputFile=history.txt --logLevel=debug  history

