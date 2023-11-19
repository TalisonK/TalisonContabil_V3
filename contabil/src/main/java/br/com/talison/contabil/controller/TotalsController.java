package br.com.talison.contabil.controller;

import br.com.talison.contabil.domain.dto.ActivityDto;
import br.com.talison.contabil.domain.dto.TotalsDto;
import br.com.talison.contabil.service.TotalsService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/totals")
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

    @PostMapping("/last-incomes")
    public ResponseEntity<List<TotalsDto>> getLastIncomes(@RequestBody TotalsDto totalsDto) {
        List<TotalsDto> retorno = totalsService.getLastIncomeTotals(totalsDto.getUserId(), totalsDto.getYear(), totalsDto.getMonth());

        return ResponseEntity.status(200).body(retorno);
    }

    @PostMapping("/last-expenses")
    public ResponseEntity<List<TotalsDto>> getLastExpenses(@RequestBody TotalsDto totalsDto) {
        List<TotalsDto> retorno = totalsService.getLastExpenseTotals(totalsDto.getUserId(), totalsDto.getYear(), totalsDto.getMonth());

        return ResponseEntity.status(200).body(retorno);
    }

    @PostMapping("/last-balances")
    public ResponseEntity<List<TotalsDto>> getLastBalances(@RequestBody TotalsDto totalsDto) {
        List<TotalsDto> retorno = totalsService.getLastBalanceTotals(totalsDto.getUserId(), totalsDto.getYear(), totalsDto.getMonth());

        return ResponseEntity.status(200).body(retorno);
    }

    @PostMapping("/income-vs-expense")
    public ResponseEntity<List<List<TotalsDto>>> getIncomeVsExpense(@RequestBody TotalsDto totalsDto) {
        List<List<TotalsDto>> retorno = totalsService.getIncomeVsExpense(totalsDto.getUserId(), totalsDto.getYear());

        return ResponseEntity.status(200).body(retorno);
    }

    @PostMapping("/timeline")
    public ResponseEntity<List<ActivityDto>> getTimeline(@RequestBody TotalsDto totalsDto) {
        List<ActivityDto> retorno = totalsService.timelineByMonth(totalsDto.getUserId(), totalsDto.getYear(), totalsDto.getMonth());

        return ResponseEntity.status(200).body(retorno);
    }

}
