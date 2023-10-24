package br.com.talison.contabil.domain.enums;

import lombok.Getter;

@Getter
public enum EnumPaymentMethod {
    CREDIT_CARD("CREDIT_CARD"),
    DEBIT_CARD("DEBIT_CARD"),
    MONEY("MONEY"),
    TRANSFER("TRANSFER"),
    PIX("PIX");

    private final String value;

    EnumPaymentMethod(String value) {
        this.value = value;
    }

}
