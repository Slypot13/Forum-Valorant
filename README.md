nso tables MYSQL : 
users
------
id
username
email
password
role
is_banned

threads
--------
id
title
content
status
user_id
created_at

messages
---------
id
content
thread_id
user_id
created_at

tags
----
id
name

thread_tags
-----------
thread_id
tag_id

reactions
----------
id
message_id
user_id
type