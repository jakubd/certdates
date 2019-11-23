[![License.md](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# certdates

Check the certificate validity period left in all your domains.  Give this script a 
text file of domains such as this:

```
https://freebsd.org
https://openbsd.org
```

and give it a threshold of the number of days of validity. For example if I am interested in 
highlighting cases where there is less than 60 days of certificate validity this value is 60.

# Compiling

You can compile this to a binary with:

```
go build certdates
```

should make a binary in the same directory

# Usage

```
./certdates --domains="domains.txt" --threshold=60
```

Will result in this type of output

TODO: put sample image here

or just 

```
./check_certs
```
