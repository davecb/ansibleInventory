# inventoryTree - alint and other ansibe inventory tools

alint reports on errors in ansible hosts files found in a specified 
directory. The directory should contain readable hosts files

A typical error message is
```
./ansible_inventory/here/hosts: AddHost() failed: host snowflake-1 exist in multiple groups: clustered, snowflake, line: 39
```
(Snowflakes are non-clustered machines, so they shouldn't be both)
