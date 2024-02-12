package br.com.talison.contabil.domain;

import br.com.talison.contabil.domain.dto.ListItemDto;
import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;

@Getter
@Setter
public class List {

    private String id;

    private String description;

    private String user;

    private LocalDateTime createdAt;

    private java.util.List<ListItemDto> list;


}
