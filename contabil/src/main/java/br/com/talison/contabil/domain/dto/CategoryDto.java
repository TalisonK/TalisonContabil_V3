package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.Setter;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import java.io.Serializable;

@Getter
@Setter
public class CategoryDto implements Serializable {

    private String id;

    @NotNull
    @NotBlank
    private String name;

    private String description;

}
