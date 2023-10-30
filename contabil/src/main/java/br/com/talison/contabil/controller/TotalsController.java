package br.com.talison.contabil.controller;

import br.com.talison.contabil.domain.dto.TotalsDto;
import br.com.talison.contabil.service.TotalsService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/totals")
@RequiredArgsConstructor
public class TotalsController {

    private final TotalsService totalsService;

    @PostMapping("/income")
    public ResponseEntity<TotalsDto> calculateIncomeTotals(@RequestBody TotalsDto totalsDto) {
        TotalsDto retorno = totalsService.getTotals(totalsDto.getYear(), totalsDto.getMonth(), totalsDto.getUserId(), "income");

        return ResponseEntity.status(200).body(retorno);
    }

    @PostMapping("/expense")
    public ResponseEntity<TotalsDto> calculateExpenseTotals(@RequestBody TotalsDto totalsDto) {
        TotalsDto retorno = totalsService.getTotals(totalsDto.getYear(), totalsDto.getMonth(), totalsDto.getUserId(), "expense");

        return ResponseEntity.status(200).body(retorno);
    }

    @PutMapping
    public ResponseEntity<TotalsDto> updateTotals(@RequestBody TotalsDto totalsDto) {
        TotalsDto retorno = totalsService.updateTotals(totalsDto.getYear(), totalsDto.getMonth(), totalsDto.getUserId(), totalsDto.getType());

        return ResponseEntity.status(200).body(retorno);
    }

}
