<p>表:&nbsp;<code>Person</code></p>

<pre>
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| email       | varchar |
+-------------+---------+
id是该表的主键列。
该表的每一行包含一封电子邮件。电子邮件将不包含大写字母。
</pre>

<p>&nbsp;</p>

<p>编写一个 SQL <strong>删除语句</strong>来 <strong>删除</strong> 所有重复的电子邮件，只保留一个id最小的唯一电子邮件。</p>

<p>以 <strong>任意顺序</strong> 返回结果表。 （<strong>注意</strong>： 仅需要写删除语句，将自动对剩余结果进行查询）</p>

<p>查询结果格式如下所示。</p>

<p>&nbsp;</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> 
Person 表:
+----+------------------+
| id | email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
<strong>输出:</strong> 
+----+------------------+
| id | email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
<strong>解释:</strong> john@example.com重复两次。我们保留最小的Id = 1。</pre>
<div><div>Related Topics</div><div><li>数据库</li></div></div><br><div><li>👍 606</li><li>👎 0</li></div>