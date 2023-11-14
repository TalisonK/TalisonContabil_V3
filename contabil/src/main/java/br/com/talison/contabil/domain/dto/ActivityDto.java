package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.Date;


@Getter
@Setter
public class ActivityDto implements Serializable {

    private String id;

    private String description;

    private String method;

    private String type;

    private String userId;

    private Double value;

    private LocalDateTime activityDate;

}
