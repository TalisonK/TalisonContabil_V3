package br.com.talison.contabil.domain.dto;

import br.com.talison.contabil.domain.User;
import lombok.Getter;
import lombok.Setter;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.Date;

@Getter
@Setter
public class IncomeDto implements Serializable {

    private String id;

    @NotNull
    @NotBlank
    private String description;

    @NotNull
    @NotBlank
    private Double value;

    @NotNull
    @NotBlank
    private String user;

    private LocalDateTime createdAt;

    @NotNull
    @NotBlank
    private Date receivedAt;

}
