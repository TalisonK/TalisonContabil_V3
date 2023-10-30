package br.com.talison.contabil.domain;

import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.io.Serializable;
import java.util.Date;

@Getter
@Setter
@Document(collection = "totals")
public class Totals implements Serializable {

    @Id
    private String id;

    @Field("value")
    private Double value;

    @Field("user")
    private String userId;

    @Field("type")
    private String type;

    @Field("createdAt")
    private Date createdAt;

    @Field("updatedAt")
    private Date updatedAt;

    @Field("year")
    private String year;

    @Field("month")
    private String month;

}
