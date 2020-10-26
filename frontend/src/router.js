  
import Vue from "vue"
import Router from "vue-router"

Vue.use(Router)

export default new Router({
  mode: "history", 
  routes:[
      {
        path: "/",
        alias: "/rosters",
        name: "rosters",
        component: () => <div>112233</div>,
      }
  ],
})