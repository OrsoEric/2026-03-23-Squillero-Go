#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <chrono>
#include <cstdint>

// -----------------------------------------------------------------------------
// STRUCTURE PAIR (uint64_t)
// -----------------------------------------------------------------------------

struct St_pair {
    uint64_t n_value[2];

    std::string to_string() const {
        return std::to_string(n_value[0]) + "-" + std::to_string(n_value[1]);
    }
};

// -----------------------------------------------------------------------------
// FILE READER (line-by-line, trim, skip empty)
// -----------------------------------------------------------------------------

std::vector<std::string> Fn_sequence_reader(const std::string& path) {
    std::ifstream file(path);
    if (!file) {
        throw std::runtime_error("Cannot open file: " + path);
    }

    std::vector<std::string> out;
    std::string line;

    while (std::getline(file, line)) {
        auto b = line.find_first_not_of(" \t\r\n");
        if (b == std::string::npos) continue;
        auto e = line.find_last_not_of(" \t\r\n");
        out.push_back(line.substr(b, e - b + 1));
    }

    return out;
}

// -----------------------------------------------------------------------------
// PARSE PAIRS (uint64_t)
// -----------------------------------------------------------------------------

std::vector<St_pair> Fn_find_pair_in_content(const std::vector<std::string>& lines) {
    std::vector<St_pair> out;

    for (const auto& line : lines) {
        std::stringstream ss(line);
        std::string chunk;

        while (std::getline(ss, chunk, ',')) {
            auto b = chunk.find_first_not_of(" \t\r\n");
            if (b == std::string::npos) continue;
            auto e = chunk.find_last_not_of(" \t\r\n");
            chunk = chunk.substr(b, e - b + 1);

            auto dash = chunk.find('-');
            if (dash == std::string::npos) {
                throw std::runtime_error("invalid pair format: " + chunk);
            }

            uint64_t left  = std::stoull(chunk.substr(0, dash));
            uint64_t right = std::stoull(chunk.substr(dash + 1));

            out.push_back({ left, right });
        }
    }

    return out;
}

// -----------------------------------------------------------------------------
// CHECK MIRROR (uint64_t)
// -----------------------------------------------------------------------------

bool Fn_is_half_number_same(uint64_t value) {
    std::string s = std::to_string(value);

    if (s.size() % 2 != 0)
        return false;

    size_t mid = s.size() / 2;
    return (s.substr(0, mid) == s.substr(mid));
}

// -----------------------------------------------------------------------------
// SERIAL SCAN WITH PROFILING
// -----------------------------------------------------------------------------

void Fn_scan_pair(
    const std::vector<St_pair>& pairs,
    uint64_t& out_num_invalid,
    uint64_t& out_sum_invalid)
{
    auto t0 = std::chrono::high_resolution_clock::now();

    uint64_t num_invalid = 0;
    uint64_t sum_invalid = 0;

    for (const auto& p : pairs) {
        for (uint64_t n = p.n_value[0]; n <= p.n_value[1]; ++n) {
            if (Fn_is_half_number_same(n)) {
                num_invalid++;
                sum_invalid += n;
            }
        }
    }

    auto t1 = std::chrono::high_resolution_clock::now();
    auto micros = std::chrono::duration_cast<std::chrono::microseconds>(t1 - t0).count();

    std::cout << "Elapsed (serial scan) [µs]: " << micros << "\n";

    out_num_invalid = num_invalid;
    out_sum_invalid = sum_invalid;
}

// -----------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------

int main() {
    try {
        std::cout << "Shaka, when the walls fell.\n";

        std::string path = "puzzle_input.txt";

        auto lines = Fn_sequence_reader(path);
        auto pairs = Fn_find_pair_in_content(lines);

        std::cout << "Pairs: " << pairs.size() << "\n";

        uint64_t num_invalid = 0;
        uint64_t sum_invalid = 0;

        Fn_scan_pair(pairs, num_invalid, sum_invalid);

        std::cout << "Num Invalid: " << num_invalid << "\n";
        std::cout << "Sum Invalid: " << sum_invalid << "\n";

    } catch (const std::exception& ex) {
        std::cerr << "Error: " << ex.what() << "\n";
        return 1;
    }

    return 0;
}
