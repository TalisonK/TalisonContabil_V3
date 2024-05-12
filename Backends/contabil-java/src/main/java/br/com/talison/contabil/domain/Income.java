package br.com.talison.contabil.domain;

import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.Date;

@Getter
@Setter
@Document(collection = "income")
public class Income implements Serializable {

    @Id
    private String id;

    @Field("description")
    private String description;

    @Field("value")
    private Double value;

    @Field("user")
    private User user;

    @Field("createdAt")
    private LocalDateTime createdAt;

    @Field("receivedAt")
    private LocalDateTime receivedAt;

    public Income(String description, Double value, User user, LocalDateTime receivedAt) {
        this.description = description;
        this.value = value;
        this.user = user;
        this.receivedAt = receivedAt;
        this.createdAt = LocalDateTime.now();
    }
}
