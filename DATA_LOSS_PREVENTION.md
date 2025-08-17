# Data Loss Prevention Checklist

## 🚨 IMMEDIATE ACTIONS COMPLETED
- [x] Renamed dangerous SQL file: `layanan_lansia.sql` → `layanan_lansia.sql.backup`
- [x] Verified ORM config is safe (no auto-migration)
- [x] Confirmed all DELETE operations use WHERE clauses
- [x] Created GORM Safety Audit Report

## 📋 DAILY MONITORING CHECKLIST

### Before Starting Work:
1. **Check database connection:**
   ```bash
   mysql -u superadmin -p layanan_lansia -e "SELECT COUNT(*) FROM auth; SELECT COUNT(*) FROM user; SELECT COUNT(*) FROM lansia;"
   ```

2. **Verify no dangerous SQL files:**
   ```bash
   find /Users/iam/Documents/Coding/svc-llt-golang -name "*.sql" -exec grep -l "DROP\|TRUNCATE\|DELETE FROM" {} \;
   ```

3. **Check DBeaver connections:**
   ```bash
   lsof -i :3306 | grep dbeaver
   ```

### If Data Loss Occurs Again:
1. **Check MySQL general log:**
   ```sql
   SHOW VARIABLES LIKE 'general_log%';
   SET GLOBAL general_log = 'ON';
   -- Monitor: /opt/homebrew/var/mysql/*.log
   ```

2. **Check recent SQL commands:**
   ```bash
   tail -f /opt/homebrew/var/mysql/iams-MacBook-Pro.local.log
   ```

3. **Check application logs:**
   ```bash
   grep -i "drop\|delete\|truncate" /path/to/app/logs/*
   ```

## 🔒 RECOMMENDED SAFETY MEASURES

### 1. Enable MySQL Binary Logging
```sql
-- Add to MySQL config (/opt/homebrew/etc/my.cnf):
[mysqld]
log-bin=mysql-bin
binlog-format=ROW
expire_logs_days=7
```

### 2. Create Database Backup Script
```bash
#!/bin/bash
# backup_db.sh
DATE=$(date +%Y%m%d_%H%M%S)
mysqldump -u superadmin -p layanan_lansia > "backup_${DATE}.sql"
echo "Backup created: backup_${DATE}.sql"
```

### 3. DBeaver Safety Settings
- Disable "Execute SQL immediately"
- Enable "Confirm dangerous SQL operations"
- Set connection to read-only mode for development

## 🎯 MONITORING POINTS

If data loss happens again, check:
1. **DBeaver query history** - Look for DROP/TRUNCATE commands
2. **MySQL general log** - All executed SQL commands
3. **Application startup logs** - Any migration or initialization
4. **File modification times** - Check if .sql files were modified
5. **Process list** - Check what processes are connected to MySQL

## 📞 ESCALATION

If data loss continues:
1. Enable MySQL general logging immediately
2. Check all database tools (DBeaver, MySQL Workbench, etc.)
3. Review any scheduled tasks or automation
4. Consider using separate development database

**Remember:** The audit confirmed ORM is safe. Focus on external tools and SQL files as the primary suspects.