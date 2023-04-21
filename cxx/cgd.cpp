#include "girlboss.hpp"
#include <fmt/core.h>
#include <filesystem>

namespace cgd {
    const char* version = "1.0.0 (c++)";
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
        fmt::print("{}: {}\n", "cgd", cgd::version);
        return 0;
    }

    if (!*sortingName || !*releaseName)
    {
        Girlboss::print_usage();
        return -1;
    }

    try {
        std::filesystem::create_directory(*sortingName);
        fmt::print("Created: {}\n", *sortingName);
        std::filesystem::current_path(*sortingName);
        std::filesystem::create_directory(*releaseName);
        std::filesystem::current_path(*releaseName);
        fmt::print("\tCreated: {}\n", *releaseName);

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
            fmt::print("\t\tCreated: {}\n", items[i]);
        }
    }

    catch (const std::exception& error)
    {
        fmt::print("{}", error.what());
        return -1;
    }

    return 0;
}
