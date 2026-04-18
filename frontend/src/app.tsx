import { MetaProvider } from "@solidjs/meta";
import { Router, useLocation } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense } from "solid-js";
import Nav from "~/components/nav/Nav";
import Title from "~/components/Title";
import "./app.css";
import Footer from "~/components/Footer";

export default function App() {
	return (
		<Router
			root={(props) => {
				const location = useLocation();
				const isBook = () => location.pathname.startsWith("/book");

				return (
					<MetaProvider>
						<Title>Meta</Title>
						<div class="min-h-screen bg-gray-50 text-gray-800 text-sm">
							<header
								class={`top-0 z-20 bg-white shadow px-6.5 h-14 flex items-center justify-between
                ${isBook() ? "relative" : "sticky"}`}
							>
								<div class="flex items-center">
									<a
										href="/"
										class="inline-block w-8 h-8 bg-[url('/favicon.png')] bg-contain bg-no-repeat bg-center"
									>
										<span class="sr-only">主页</span>
									</a>
								</div>
								<div class="absolute left-1/2 -translate-x-1/2">
									<Nav />
								</div>
								<div class="flex items-center"></div>
							</header>
							<Suspense>{props.children}</Suspense>
							<Footer />
						</div>
					</MetaProvider>
				);
			}}
		>
			<FileRoutes />
		</Router>
	);
}
