import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
//import "./assets/tailwind.css";

import PrimeVue from "primevue/config";
import "primevue/resources/primevue.min.css";
import "primevue/resources/themes/bootstrap4-dark-purple/theme.css";

import Button from "primevue/button";

import DataTable from "primevue/datatable";
import Column from "primevue/column";
import ColumnGroup from "primevue/columngroup"; // optional
import Row from "primevue/row"; // optional

import SelectButton from "primevue/selectbutton";

const app = createApp(App);
app.use(router);
app.use(PrimeVue);
app.component("Button", Button);
app.component("DataTable", DataTable);
app.component("Column", Column);
app.component("ColumnGroup", ColumnGroup);
app.component("Row", Row);
app.component("SelectButton", SelectButton);

app.mount("#app");
