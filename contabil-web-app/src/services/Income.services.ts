import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";


@Injectable()

export class IncomeServices{

    constructor(private http: HttpClient) {
    }

    url = 'http://localhost:8080/income';

    createIncome(income: any) {
        return this.http.post(this.url, income);
    }

}
