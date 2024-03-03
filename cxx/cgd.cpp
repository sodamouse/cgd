#define GIRLBOSS_IMPL
#include "girlboss.hpp"

#include <filesystem>
#include <iostream>

namespace cgd {

constexpr char* VERSION = "CGD 1.1.0 (c++)";

} // namespace cgd

int main(int argc, char* argv[])
{
    const char** sortingName = Girlboss::flag_str("-a", nullptr, "Game sorting name");
    const char** releaseName = Girlboss::flag_str("-b", nullptr, "Game release name");
    bool* showHelp = Girlboss::flag_option("-h", false, "Show this message");
    bool* showVersion = Girlboss::flag_option("-v", false, "Display program version information");
    Girlboss::parse("cgd", argc, argv);

    if (*showHelp)
    {
        Girlboss::print_usage();
        return 0;
    }

    if (*showVersion)
    {
        std::cout << cgd::VERSION << '\n';
        return 0;
    }

    if (!*sortingName || !*releaseName)
    {
        Girlboss::print_usage();
        return -1;
    }

    try {
        std::filesystem::create_directory(*sortingName);
        std::cout << "Created: " << *sortingName << '\n';
        std::filesystem::current_path(*sortingName);
        std::filesystem::create_directory(*releaseName);
        std::filesystem::current_path(*releaseName);
        std::cout << "\tCreated: " << *releaseName << '\n';

        const char* items[5] = {
            "dlc",
            "extras",
            "instructions",
            "setup",
            "updates"
        };

        for (int i = 0; i < 5; ++i)
        {
            std::filesystem::create_directory(items[i]);
            std::cout << "\t\tCreated: " << items[i] << '\n';
        }
    }

    catch (const std::exception& error)
    {
        std::cout << error.what() << '\n';
        return -1;
    }

    return 0;
}
