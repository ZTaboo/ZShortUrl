import {createRouter, createWebHashHistory, RouteComponent} from "vue-router";

interface routerCon {
    path: string,
    component: RouteComponent
}

const routes: routerCon[] = [
    {
        path: '/',
        component: () => import('../views/Home.vue'),
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router;