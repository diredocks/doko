import { navLinks } from "./navLinks";

const baseLinkClass =
	"relative flex items-center px-4 " +
	"after:content-[''] after:absolute after:left-1/2 after:-translate-x-1/2 after:bottom-1 " +
	"after:h-[2px] after:w-[50%] after:bg-gray-800 " +
	"after:scale-x-0 after:origin-center " +
	"hover:after:scale-x-100 after:transition-transform";

export default function Nav() {
	return (
		<nav class="h-14">
			<ul class="flex h-full">
				{navLinks.map((link) => (
					<li
						class={`flex ${
							link.special
								? "bg-[url(/special-nav.png)] bg-no-repeat bg-center bg-cover"
								: ""
						}`}
					>
						<a href={link.href} class={baseLinkClass}>
							{link.label}
						</a>
					</li>
				))}
			</ul>
		</nav>
	);
}
