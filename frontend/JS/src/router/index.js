import { createRouter, createWebHashHistory } from "vue-router";
import Home from "@/views/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: function () {
      return import(/* webpackChunkName: "about" */ "../views/About.vue");
    },
  },
  {
    path: "/newConnection",
    name: "newConnection",
    component: function () {
      return import(/* webpackChunkName: "about" */ "../views/NewConnection.vue");
    },
  },
  {
    path:"/home/welcome",
    name:"welcome",
    component:function (){
      return import(/* webpackChunkName: "about" */ "../views/Content_welcome.vue")
    }
  },
  // {
  //   path:"/:key",
  //   name:"welcome",
  //   children:[
  //     {
  //       path: "",
  //       component: welcome,
  //     }
  //   ]
  // },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
