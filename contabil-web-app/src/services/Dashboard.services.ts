import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Totals} from "../domains/Totals";
import {Injectable} from "@angular/core";


@Injectable()
export class DashboardServices {

    url = `http://localhost:8080/totals`;
    constructor(private http: HttpClient) {}


    getIncomes(userId: string, year: string, month: string): Observable<Totals[]> {
        return this.http.post<Totals[]>(`${this.url}/last-incomes`, {userId, year, month});
    }

    getExpense(userId: string, year: string, month: string): Observable<Totals[]> {
        return this.http.post<Totals[]>(`${this.url}/last-expenses`, {userId, year, month});
    }
}
