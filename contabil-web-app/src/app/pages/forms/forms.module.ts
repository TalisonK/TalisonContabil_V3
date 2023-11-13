import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PrimengModule} from "../../shared/primeng.module";
import {CategoryServices} from "../../../services/category.services";
import {MethodsUtil} from "../../shared/utils/methods.util";
import {ExpenseServices} from "../../../services/expense.services";
import {IncomeServices} from "../../../services/income.services";


@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        PrimengModule

    ],
    providers: [
        CategoryServices,
        MethodsUtil,
        ExpenseServices,
        IncomeServices
    ]
})
export class FormsModule {
}
