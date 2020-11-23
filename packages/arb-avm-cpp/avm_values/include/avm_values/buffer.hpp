/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef buffer_hpp
#define buffer_hpp

#include <memory>
#include <vector>
#include <cstdint>
#include <iostream>
#include <avm_values/bigint.hpp>

const uint64_t LEAF_SIZE = 1024;
const uint64_t NODE_SIZE = 8;
const uint64_t ALIGN = LEAF_SIZE;

inline uint64_t calc_len(int h) {
    if (h == 0) {
        return LEAF_SIZE;
    }
    return NODE_SIZE*calc_len(h-1);
}

struct Packed {
    uint256_t hash;
    uint64_t size; // total size
    int packed; // packed levels
};

Packed zero_packed(uint64_t sz);

class RawBuffer {
   private:
    int level;
    bool saved;
    Packed savedHash;

    std::shared_ptr<std::vector<uint8_t> > leaf;
    std::shared_ptr<std::vector<RawBuffer> > node;

    RawBuffer(std::shared_ptr<std::vector<uint8_t> > leaf_) : leaf(leaf_), node(nullptr) {
        // std::cerr << "creating buffer 0" << std::endl;
        level = 0;
        saved = false;
    }

    RawBuffer(std::shared_ptr<std::vector<RawBuffer> > node_, int level_) : leaf(nullptr), node(node_) {
        // std::cerr << "creating buffer " << level_ << std::endl;
        level = level_;
        saved = false;
    }

    RawBuffer(int level_, bool) : leaf(nullptr), node(nullptr) {
        // std::cerr << "creating buffer " << level_ << std::endl;
        level = level_;
        saved = true;
        savedHash = zero_packed(calc_len(level));
    }

   public:
    RawBuffer() : leaf(nullptr), node(nullptr) {
        // std::cerr << "creating buffer\n";
        level = 0;
        saved = true;
        savedHash = zero_packed(LEAF_SIZE);
    }

    RawBuffer set(uint64_t offset, uint8_t v) {
        std::vector<uint8_t> arr(1);
        arr[0] = v;
        return set_many(offset, arr);
    }

    // Note: pos and len must be aligned so that the data to be written is in one leaf
    RawBuffer set_many(uint64_t offset, std::vector<uint8_t> &arr) {
        // std::cerr << "setting buffer " << level << " at " << offset << " to " << std::hex << int(v) << std::endl;
        if (level == 0) {
            if (offset >= LEAF_SIZE) {
                std::shared_ptr<std::vector<uint8_t> > empty = std::make_shared<std::vector<uint8_t>>();
                std::shared_ptr<std::vector<RawBuffer> > vec = std::make_shared<std::vector<RawBuffer>>();
                vec->push_back(RawBuffer(leaf));
                for (uint64_t i = 1; i < NODE_SIZE; i++) {
                    vec->push_back(RawBuffer(empty));
                }
                RawBuffer buf = RawBuffer(vec, 1);
                return buf.set_many(offset, arr);
            }
            auto buf = leaf ? std::make_shared<std::vector<uint8_t> >(*leaf) : std::make_shared<std::vector<uint8_t> >();
            if (buf->size() < LEAF_SIZE) {
                // std::cerr << "resize buf" << std::endl;
                buf->resize(LEAF_SIZE, 0);
            }
            for (unsigned int i = 0; i < arr.size(); i++) {
                (*buf)[offset+i] = arr[i];
            }
            // std::cerr << "updated leaf " << level << " at " << offset << " to " << std::hex << int(v) << std::endl;
            return RawBuffer(buf);
        } else {
            if (offset >= calc_len(level)) {
                std::shared_ptr<std::vector<RawBuffer> > vec = std::make_shared<std::vector<RawBuffer>>();
                vec->push_back(RawBuffer(node, level));
                for (uint64_t i = 1; i < NODE_SIZE; i++) {
                    vec->push_back(RawBuffer(level, true));
                }
                RawBuffer buf = RawBuffer(vec, level+1);
                return buf.set_many(offset, arr);
            }
            auto vec = std::make_shared<std::vector<RawBuffer> >(node ? *node : RawBuffer::make_empty(level-1));
            auto cell_len = calc_len(level-1);
            (*vec)[offset / cell_len] = (*vec)[offset / cell_len].set_many(offset % cell_len, arr);
            return RawBuffer(vec, level);
        }
    }

    static std::vector<RawBuffer> make_empty(int level) {
        auto vec = std::vector<RawBuffer>();
        for (uint64_t i = 0; i < NODE_SIZE; i++) {
            vec.push_back(RawBuffer(level, true));
        }
        return vec;
    }

    uint8_t get(uint64_t pos) const {
        auto res = get_many(pos, 1);
        return res[0];
    }

    // Note: pos and len must be aligned so that the data to be read is from one leaf
    std::vector<uint8_t> get_many(uint64_t pos, int len) const {
        if (level == 0) {
            auto res = std::vector<uint8_t>(len, 0);
            for (int i = 0; i < len; i++) {
                if (!leaf) res[i] = 0;
                else if (leaf->size() < pos+i) res[i] = 0;
                else res[i] = (*leaf)[pos+i];
            }
            return res;
        } else {
            uint64_t ln = calc_len(level);
            uint64_t cell_len = calc_len(level-1);
            if (pos > ln || !node) {
                return std::vector<uint8_t>(len, 0);
            }
            return (*node)[pos / cell_len].get_many(pos % cell_len, len);
        }
    }

    uint256_t hash();
    Packed hash_aux();

    void serialize(std::vector<unsigned char>& value_vector);

    RawBuffer normalize();

    uint64_t size() const {
        return calc_len(level);
    }

    static RawBuffer deserialize(const char *buf, int level, int &len);

    friend class Buffer;

};

class Buffer {
   private:
    std::shared_ptr<RawBuffer> buf;

    Buffer(const RawBuffer &buffer) {
        buf = std::make_shared<RawBuffer>(buffer);
    }

   public:
    Buffer() {
        buf = std::make_shared<RawBuffer>();
    }

    Buffer set(uint64_t offset, uint8_t v) {
        return Buffer(buf->set(offset, v));
    }

    Buffer set_many(uint64_t offset, std::vector<uint8_t> arr) {
        return Buffer(buf->set_many(offset, arr));
    }

    uint8_t get(uint64_t pos) const {
        return buf->get(pos);
    }

    std::vector<uint8_t> get_many(uint64_t pos, int len) const {
        return buf->get_many(pos, len);
    }

    uint64_t size() const {
        return buf->size();
    }

    uint256_t hash() const {
        return buf->hash();
    }

    void serialize(std::vector<unsigned char>& value_vector) const {
        RawBuffer nbuf = buf->normalize();
        value_vector.push_back(static_cast<uint8_t>(nbuf.level));
        nbuf.serialize(value_vector);
    }

    static Buffer deserialize(const char *buf, int &len) {
        uint8_t level = buf[0];
        len++;
        return Buffer(RawBuffer::deserialize(buf+1, level, len));
    }
};

inline uint256_t hash(const Buffer& b) {
    return b.hash();
}


inline bool operator==(const Buffer& val1, const Buffer& val2) {
    return hash(val1) == hash(val2);
}

#endif /* buffer_hpp */
