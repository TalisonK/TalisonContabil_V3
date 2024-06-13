package br.com.talison.contabil.domain.dto;

import br.com.talison.contabil.domain.User;
import lombok.Getter;
import lombok.Setter;
import org.springframework.data.mongodb.core.mapping.Field;

import java.io.Serializable;
import java.util.Date;

@Getter
@Setter
public class TotalsDto implements Serializable {

    private String id;

    private Double value;

    private String userId;

    private String type;

    private Date createdAt;

    private Date updatedAt;

    private String year;

    private String month;

}
