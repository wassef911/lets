## 1. Searching Files & Text
Replace: grep, find, and basic awk

```sh
# Find files containing "error" (case-insensitive)
# grep -r -i "error" /var/log
$ lets search files for "error" in /var/log	

# Count occurrences of "404" in a file 
# grep -c "404" access.log
$ lets count matches "404" in access.log	

# Find logs older than a week	
# find /home -name "*.log" -mtime +7
$ lets find files named "*.log" in /home older than 7 days	
```

## 2. Disk & Storage:
Replace: df, du

```sh
# Display disk usage for all mounts	
# df -h
$ lets show disk space	

# Display directory disk usage
# sizedu -sh /var
$ lets show folder size for /var

# Find large files	
# find /home -type f -size +100M
$ lets show files over 100MB in /home	
```

## 3. Process Management:
Replace: top, htop, ps, kill

```sh
# List all running processes	
# ps aux
$ lets show processes	

# Live system resource view	
# htop
$ lets show resources	

# Terminate by process name	
# pkill apache2
$ lets kill process "apache2"	
```

## 4. Data Extraction:
Replace: awk, cut, sed

```sh
# Print the 3rd column	
# awk '{print $3}' data.csv
$ lets get column 3 from data.csv	

# In-place text replacement	
# sed -i 's/foo/bar/g' file.txt
$ lets replace "foo" with "bar" in file.txt	
```


# TODO
* manuall test all commands
    - search done
* handle errors and panic properly 
* write tests for all packages
* getting files other than CSV
* better validation
* structured output / better logger