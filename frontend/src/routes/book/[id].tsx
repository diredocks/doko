import { Icon } from "@iconify-icon/solid";
import arrow from "@iconify-icons/mdi/keyboard-arrow-right";
import help from "@iconify-icons/mdi/people-group";
import { useParams } from "@solidjs/router";
import Title from "~/components/Title";

export default function Book() {
	const params = useParams();
	const tags = ["心理学", "教育", "社会", "刑侦", "犯罪", "纪实", "非虚构"];

	return (
		<>
			<div class="sticky z-10 top-0 bg-gray-100 text-xs">
				<div class="w-200 mx-auto">
					<div class="items-center flex h-11.25">
						<a href="/">分类</a>
						<span class="w-2.5 mx-2 inline-flex items-center">
							<Icon icon={arrow} />
						</span>
						<a href="/">非虚构</a>
						<span class="w-2.5 mx-2 inline-flex items-center">
							<Icon icon={arrow} />
						</span>
						<a href="/">人性的深渊：吴谢宇案</a>
					</div>
				</div>
			</div>
			<main class="w-200 mx-auto mb-10 mt-8.75">
				<Title>详情</Title>
				<div class="flex">
					<div class="flex-1">
						<div class="flex border-b border-gray-200 pb-5 mb-5">
							<div class="mr-5">
								<img
									class="w-35 aspect-3/4 object-cover shadow"
									alt="cover"
									src="https://pic.arkread.com/cover/ebook/f/713988758.1775548446.jpg"
								/>
							</div>
							<div class="grow text-xs">
								<h1 class="font-semibold text-xl">人性的深渊：吴谢宇案</h1>
								<div class="mt-2 leading-5.25">
									<p>
										<span class="mr-1.25 text-gray-500">作者</span>
										<span class="mr-1.25 text-gray-400">吴琪</span>
										<span class="mr-1.25 text-gray-400">王珊</span>
									</p>
									<p>
										<span class="mr-1.25 text-gray-500">类别</span>
										<span class="mr-1.25">非虚构</span>
									</p>
									<p>
										<span class="mr-1.25 text-gray-500">出版商</span>
										<span class="mr-1.25">生活·读书·新知三联书店</span>
									</p>
									<p>
										<span class="mr-1.25 text-gray-500">出版时间</span>
										<span class="mr-1.25">2025</span>
									</p>
									<p>
										<span class="mr-1.25 text-gray-500">上传者</span>
										<span class="mr-1.25 text-gray-400">管理员</span>
									</p>
									<p>
										<span class="mr-1.25 text-gray-500">ISBN</span>
										<span class="mr-1.25 text-gray-400">9787807684473</span>
									</p>
								</div>
								<div class="mt-3 flex items-center text-gray-400">
									<span
										class="relative inline-flex items-center before:content-['★★★★★'] before:tracking-[2px] before:bg-[linear-gradient(90deg,#f5a623_calc(var(--rating)/10*100%),#ddd_calc(var(--rating)/10*100%))] before:bg-clip-text before:text-transparent"
										style="--rating:7.5"
									/>
									<span class="ml-2">7.5</span>
								</div>
								<div class="mt-5">
									<a
										class="flex font-semibold my-5 px-4 py-2 bg-[#fad36f] rounded-lg items-center gap-1.5 w-full"
										href="/"
									>
										<i class="w-4 h-4">
											<Icon height={16} width={16} icon={help} />
										</i>
										<span>上传你的电子书，在 Doko 分享和保存人类的知识</span>
									</a>
									<a
										href="/"
										class="w-38 h-10 bg-teal-400 hover:bg-teal-500 text-white rounded-full flex items-center justify-center text-base"
									>
										在线阅读
									</a>
								</div>
							</div>
						</div>
						<div class="mt-9 mb-5">
							<h3 class="font-semibold text-lg mb-3.75">简介</h3>
							<p class="leading-6">
								吴谢宇案是一起备受关注的刑事案件。2015年7月，时年21岁的北京大学学生吴谢宇在家中杀害母亲谢天琴，制造母亲陪同其出国留学的假象，骗取亲友144万元用于挥霍，购买十余张身份证件隐匿身份逃亡，直到2019年4月在重庆机场被捕。法院审理认定其作案前精心预谋，手段极其残忍，严重违背人伦道德，最终以故意杀人罪、诈骗罪、买卖身份证件罪数罪并罚。经最高人民法院核准，于2024年1月被执行死刑。
							</p>
						</div>
						<div>
							<h3 class="font-semibold text-lg mb-3.75">
								评论
								<span class="ml-1.25 text-sm text-gray-400">93</span>
							</h3>
						</div>
					</div>
					<aside class="w-64">
						<div>
							<ul>
								{tags.map((tag) => (
									<li class="inline-block text-xs">
										<a
											class="bg-gray-200 text-gray-500 rounded-full block px-2.5 py-1.5 ml-1.75 mb-1.5 leading-none"
											href="/"
										>
											<span>{tag}</span>
										</a>
									</li>
								))}
							</ul>
						</div>
					</aside>
				</div>
			</main>
		</>
	);
}
