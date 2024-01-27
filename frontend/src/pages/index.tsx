import { useGetUserInfo } from '../hooks/getUserInfo';
import { useGetUserFollowers } from '../hooks/getUserFollowers';
import { useGetUserRepositories } from '../hooks/getUserRepos';
import { useValidateAuth } from '../hooks/validateAuth';
import { Header, Sidebar, Section, RepoCard, Footer } from '../components';

export default function Home() {
  useGetUserInfo();
  useGetUserFollowers();
  useGetUserRepositories();
  useValidateAuth();

  return (
    <div className='min-h-screen flex flex-col justify-between'>
      <section>
        <Header />
        <div className='flex justify-center mt-6'>
          <div className='px-4 grid lg:grid-cols-4 gap-4 container'>
            {/* Sidebar section */}
            <div className='md:px-4 col-span-4 lg:col-span-1'>
              <Sidebar />
            </div>

            {/* Repositories section */}
            <div className='lg:p-6 col-span-4 lg:col-span-3 bg-white lg:border rounded-lg'>
              <Section />
              <RepoCard />
            </div>
          </div>
        </div>
      </section>

      <Footer />
    </div>
  );
}
