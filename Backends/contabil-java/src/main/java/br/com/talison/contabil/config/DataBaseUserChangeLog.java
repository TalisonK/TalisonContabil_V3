package br.com.talison.contabil.config;

import io.mongock.api.annotations.BeforeExecution;
import io.mongock.api.annotations.ChangeUnit;
import org.springframework.data.mongodb.core.MongoTemplate;

@ChangeUnit(id= "contabil", order = "001", author = "talison.kennedy")
public class DataBaseUserChangeLog {

    private final MongoTemplate mongoTemplate;

    public DataBaseUserChangeLog(MongoTemplate mongoTemplate) {
        this.mongoTemplate = mongoTemplate;
    }

    @BeforeExecution
    public void before() {

        mongoTemplate.createCollection("user");
        mongoTemplate.createCollection("expense");
        mongoTemplate.createCollection("income");
        mongoTemplate.createCollection("category");
        mongoTemplate.createCollection("totals");
    }

}
