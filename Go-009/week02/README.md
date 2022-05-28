#### 作业：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层？

```
errors.Wrapf(common.ErrNotFound, "sql: %s error: %v", sqlQuery, err)
```
