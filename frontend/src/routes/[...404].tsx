import { HttpStatusCode } from "@solidjs/start";
import Title from "~/components/Title";

export default function NotFound() {
	return (
		<main class="w-200 mx-auto mb-10 mt-7.5">
			<Title>404</Title>
			<HttpStatusCode code={404} />
			<h1>你似乎来到了没有知识存在的荒原</h1>
		</main>
	);
}
