export default function Footer() {
	return (
		<footer class="py-2.5 border-t border-gray-200">
			<div class="p-10 flex items-center justify-center text-gray-500">
				<div class="flex-col space-y-4 w-200">
					<ul class="flex space-x-4">
						<a class="hover:text-teal-500" href="/about">
							关于
						</a>
						<a class="hover:text-teal-500" href="/about">
							帮助
						</a>
						<a class="hover:text-teal-500" href="/about">
							开发者
						</a>
					</ul>
					<div class="flex text-xs text-gray-400">
						© 2026 read.doko.io, built with love by Doko contributors.
					</div>
				</div>
			</div>
		</footer>
	);
}
