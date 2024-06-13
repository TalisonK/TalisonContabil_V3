package br.com.talison.contabil.domain.enums;

import lombok.Getter;

@Getter
public enum EnumRole {
    ROLE_ADMIN("ROLE_ADMIN"),
    ROLE_USER("ROLE_USER");

    private final String value;

    EnumRole(String value) {
        this.value = value;
    }
}
