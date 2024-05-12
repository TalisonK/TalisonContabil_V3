package br.com.talison.contabil.domain;

import br.com.talison.contabil.domain.enums.EnumPaymentMethod;
import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.Date;

@Getter
@Setter
@Document(collection = "expense")
public class Expense implements Serializable {

    @Id
    private String id;

    @Field("description")
    private String description;

    @Field("paymentMethod")
    private EnumPaymentMethod paymentMethod;

    @Field("category")
    private Category category;

    @Field("value")
    private Double value;

    @Field("user")
    private User user;

    @Field("createdAt")
    @CreatedDate
    private LocalDateTime createdAt;

    @Field("paidAt")
    private LocalDateTime paidAt;

    @Field("actualParcel")
    private Integer actualParcel;

    @Field("totalParcel")
    private Integer totalParcel;

    public Expense() {
        this.createdAt = LocalDateTime.now();
    }

    public Expense(String description, EnumPaymentMethod paymentMethod, Category category, Double value, User user, LocalDateTime paidAt) {
        this();
        this.description = description;
        this.paymentMethod = paymentMethod;
        this.category = category;
        this.value = value;
        this.user = user;
        this.paidAt = paidAt;
    }

    public Expense(String description, EnumPaymentMethod paymentMethod, Category category, Double value, User user, LocalDateTime paidAt, Integer actualParcel, Integer totalParcel) {
        this(description, paymentMethod, category, value, user, paidAt);
        this.actualParcel = actualParcel;
        this.totalParcel = totalParcel;
    }
}
