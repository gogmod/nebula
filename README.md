# nebula
### docker container test ####
    docker pull vesoft/nebula-graph
    docker run --rm -it vesoft/nebula-graph bash
    
    #启动所有的nebula服务 graphd, metad and storaged
    ./scripts/nebula.service start all
    #通过console连接数据库
    ./bin/nebula -u user -p password
    
    #列出所有注册的storage hosts
    SHOW HOSTS
    #创建一个名为sp的Space, partition_num指定一个副本中的分区数量，replica_factor 指定集群的副本数量
    CREATE SPACE sp(partition_num=1024,replica_factor=1);
    #使用sp
      USE sp;
    #定义Schema
      #CREATE TAG定义带有名字和属性列表的标签
       CREATE TAG course(name string, credits int);
       create tag building(name string);
       create tag student(name string, age int, gender string);
      #定义两个edge types
       create edge like(likeness double);
       create edge select(grade int);
      #SHOW TAGS 查看标签列表
       SHOW TAGS
      #SHOW EDGES 查看边类型列表
       SHOW EDGES
      #DESCRIBE TAG 查看标签的属性
       DESCRIBE TAG
      #DESCRIBE EDGE 查看边类型的属性
      #插入数据
      INSERT VERTEX student(name, age, gender)VALUES 200:("Monica",16,"female");
      insert vertex student(name, age, gender) VALUES 201("Gary",18,"male");
      insert vertex course(name,credits),building(name)VALUES 101:("Month",3,"No5");
      insert vertex course(name,credits),building(name)VALUEs 102:("English",6,"No11");
      #插入一条边，多条边
      insert edge select(grade)VALUES 200 -> 101:(5);
      insert edge select(grade)VALUES 200 -> 102:(3);
      insert edge select(grade)VALUES 201 -> 102:(3);
      insert edge select(grade)VALUES 202 -> 102:(3);
      insert edge like(likeness) VALUES 200 -> 201:(92.5);
      insert edge like(likeness) VALUES 201 -> 200:(85.6);
      insert edge like(likeness) VALUES 201 -> 202:(93.6);
      #查询点201喜欢的点
      GO FROM 201 OVER like;
      #查询点201喜欢的,并且年龄大于17的点，
      GO FROM 201 OVER like where $$.student.age >= 17 YIELD $$.student.name as Friend, $$.student.age as Age, $$.student.gender AS Gender;
      $a=GO FROM 201 OVER like yield like._dst as id; GO FROM $a.id OVER select YIELD $^.student.name AS Student, $$.course.name AS Course, select.grade AS Grade

      
  
    

