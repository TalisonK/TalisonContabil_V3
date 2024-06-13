package br.com.talison.contabil.domain.dto;

import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;

@Getter
@Setter
public class ListItemDto implements Serializable {

    private String name;

    private String price;

}
