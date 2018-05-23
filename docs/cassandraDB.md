# Cassandra Schema

## Creation Scripts

### Create the name space

`CREATE KEYSPACE api_victorsesma WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`

### Create the table for the curriculum vitae

```cql
CREATE TABLE api_victorsesma.curriculum_vitae (
    section_type text,
    show_order int,
    start_date date,
    end_date date,
    name text,
    summary text,
    description text,
    PRIMARY KEY(section_type,show_order)
);
```

### Insert the data

```cql
INSERT INTO api_victorsesma.curriculum_vitae
(section_type, show_order, start_date,end_date,name,summary,description)
VALUES
(
    'work_experience',
    1,
    '2016-06-05',
    '2018-05-14',
    'Full Stack Developer and Support Leader',
    'Full Stack Developer and Support Leader in leading industry SmarterClick.com. (London)',
    'As an small company in continuous grow, Smarter Click has challenged me from the beginning in different areas inside the tech and development world.
- I have been designing the MySQL tables and wrote the backend code that stores all the transactions coming from a 3rd party REST API using PHP.
- I have written all the JavaSript that shows in our admin panel the sensitive transaction stats data, including graphs, using Json as as method for loading all client-specific data.
- I have been challenged by the grow in traffic needing to use cache systems such as Redis.
- I have been involved in security upgrades, like Multifactor Auth, improving the restrictions of data each client can see, AWS security groups...
- I have been using continuously AWS, including sys admin (Apache, SSL certificates...) of admin machines and upgrades of Auto Scaling Groups, Rote53, SES, S3 and other products.
- I have been challenged by the stress of peak periods (Black Friday, Christmas...) with a big amount of work searching quick and efficient solutions for our clients writing custom JavaScript code that fixes specific issues.
- Apart of the tech specific challenges,'
);

INSERT INTO api_victorsesma.curriculum_vitae
(section_type, show_order, start_date,end_date,name,summary,description)
VALUES
(
    'work_experience',
    2,
    '2016-01-01',
    '2016-04-01',
    'Project Manager',
    'Project Manager for the IT department in WatchFit. Managing a team of 4 people: 3 remote workers
and 1 intern in the office. (London)',
    '- Project Management: Using Agile Methodology approach and RedMine as a project management
software tool.
- Testing: Ensuring that the software is ready to be deployed.
- GitHub: Manage software versions with GigHub. Experience with its functionalities (branches, merges, checkouts, pull, push...)
- Coding: front-end pages (HTHML/CSS) and PHP including a bit of Yii framework.
- WordPress: Including fixing problems, changes, creating plug-ins for emails notification, Disqus integration...
- SQL: Design for storing new features needs and SQL sentences for getting information.
- Security: Copies policies, etc.
- Servers Administration: We have been using Amazon Web Services and Digital Ocean. Some AWS features I feel comfortable using are:
- EC2: creating new instances, elastic IPs, security groups...
- S3: For storing information.
- Route 53: For managing the DNS'
);

INSERT INTO api_victorsesma.curriculum_vitae
(section_type, show_order, start_date,end_date,name,summary,description)
VALUES
(
    'work_experience',
    3,
    '2015-05-01',
    '2015-12-01',
    'WatchFit CTO Assistant',
    'Project Manager Assistant, System Administrator, PHP developer, Tester in WatchFit.com (London)',
    'During this period I did training in all the task explained above.'
);

INSERT INTO api_victorsesma.curriculum_vitae
(section_type, show_order, start_date,end_date,name,summary,description)
VALUES
(
    'work_experience',
    4,
    '2015-07-01',
    '2015-09-01',
    'IT Manager in Performing Production',
    'Performing Production (London)',
    'IT manager in Performing Production. I re-built the WordPress company page using a template and customised it.
I fixed problems with the computers: re-installing Windows and Ubuntu, setting-up printers...'
);

INSERT INTO api_victorsesma.curriculum_vitae
(section_type, show_order, start_date,end_date,name,summary,description)
VALUES
(
    'work_experience',
    5,
    '2015-05-01',
    '2016-07-01',
    'PHP Intern in Lmental.com',
    'PHP Coding Intern in Lmental.com (Alicante, Spain.) Internship where I learned PHP, WordPress, web development, javascript/jQuery, jSon, bootstrap,',
    'SEO, rest APIS, and other web technologies.'
);
```