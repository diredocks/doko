import Title from "~/components/Title";

const categories = [
	"小说",
	"历史",
	"文学",
	"社科",
	"经管",
	"非虚构",
	"设计",
	"生活",
	"科技",
	"少儿",
];

const trending = [
	{
		id: 0,
		title: "1984",
		author: "乔治·奥威尔",
		cover: "https://pic.arkread.com/cover/ebook/f/9553470.1653698690.jpg",
	},
	{
		id: 1,
		title: "高效学习：曹将的公开课",
		author: "曹将",
		cover: "https://pic.arkread.com/cover/ebook/f/700982869.1772790277.jpg",
	},
	{
		id: 2,
		title: "人性的深渊：吴谢宇案",
		author: "吴琪",
		cover: "https://pic.arkread.com/cover/ebook/f/713988758.1775548446.jpg",
	},
	{
		id: 3,
		title: "历史中的大与小",
		author: "马伯庸",
		cover: "https://pic.arkread.com/cover/ebook/f/714980027.1775722184.jpg",
	},
	{
		id: 4,
		title: "中国的奋斗：1600—2000",
		author: "徐中约",
		cover: "https://pic.arkread.com/cover/ebook/f/707996974.1774340790.jpg",
	},
	{
		id: 5,
		title: "羊脂球",
		author: "莫泊桑",
		cover: "https://pic.arkread.com/cover/ebook/f/715391044.1775803949.jpg",
	},
];

export default function Home() {
	return (
		<main class="w-200 mx-auto mb-10 mt-7.5">
			<Title>主页</Title>
			<div class="flex">
				<aside class="w-48 pr-10">
					<section>
						<h1 class="text-lg font-medium px-4 pb-2">分类</h1>
						<ul>
							{categories.map((category) => (
								<li>
									<a
										class="px-4 py-2 block hover:bg-black/5 transition-colors"
										href="/"
									>
										{category}
									</a>
								</li>
							))}
						</ul>
					</section>
				</aside>
				<div class="flex-1">
					<section class="pb-6">
						<h1 class="text-lg border-b font-medium border-gray-300 pb-2 mb-4">
							时下流行
						</h1>
						<div class="grid grid-cols-4 gap-6">
							{trending.map((book) => (
								<div class="flex flex-col">
									<a href={`/book/${book.id}`}>
										<img
											src={book.cover}
											alt={book.title}
											class="w-full aspect-3/4 object-cover shadow transform transition-transform duration-300 hover:scale-105 hover:-rotate-2"
										/>
										<div class="mt-2 mb-0.5 line-clamp-2">{book.title}</div>
									</a>
									<div class="text-xs text-gray-500">{book.author}</div>
								</div>
							))}
						</div>
					</section>
				</div>
			</div>
		</main>
	);
}
