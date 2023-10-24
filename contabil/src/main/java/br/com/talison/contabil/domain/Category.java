package br.com.talison.contabil.domain;


import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.io.Serializable;

@Getter
@Setter
@Document(collection = "category")
public class Category implements Serializable {

    @Id
    private String id;

    @Field("name")
    @Indexed(unique = true)
    private String name;

    @Field("description")
    private String description;

    public Category() {
    }

    public Category(String name) {
        this.name = name;
        this.description = "";
    }

    public Category(String name, String description) {
        this(name);
        this.description = description;
    }
}
