import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Totals} from "../domains/Totals.interface";
import {Injectable} from "@angular/core";
import {environment} from "../environments/environment";
import {DashboardInterface} from "../domains/dashboard.interface";


@Injectable()
export class DashboardServices {

    url = environment.API_BASE_URL + '/totals';
    constructor(private http: HttpClient) {}

    getDashboard(userId: string, year: string, month: string): Observable<DashboardInterface> {
        return this.http.post<DashboardInterface>(`${this.url}/dashboard`, {userId, year, month});
    }
}
